import { useEffect } from "react";
import {
  fetchChairPostCoordinate,
  fetchChairPostRideStatus,
} from "~/api/api-components";
import { RideId } from "~/api/api-parameters";
import { Coordinate } from "~/api/api-schemas";
import { useSimulatorContext } from "~/contexts/simulator-context";
import {
  setSimulatorCurrentCoordinate,
  setSimulatorStartCoordinate,
} from "~/utils/storage";

const move = (
  currentCoordinate: Coordinate,
  targetCoordinate: Coordinate,
): Coordinate => {
  switch (true) {
    case currentCoordinate.latitude !== targetCoordinate.latitude: {
      const sign =
        targetCoordinate.latitude - currentCoordinate.latitude > 0 ? 1 : -1;
      return {
        latitude: currentCoordinate.latitude + sign * 1,
        longitude: currentCoordinate.longitude,
      };
    }
    case currentCoordinate.longitude !== targetCoordinate.longitude: {
      const sign =
        targetCoordinate.longitude - currentCoordinate.longitude > 0 ? 1 : -1;
      return {
        latitude: currentCoordinate.latitude,
        longitude: currentCoordinate.longitude + sign * 1,
      };
    }
    default:
      throw Error("Error: Expected status to be 'Arraived'.");
  }
};

const currentCoodinatePost = (coordinate: Coordinate) => {
  setSimulatorCurrentCoordinate(coordinate);
  void fetchChairPostCoordinate({
    body: coordinate,
  }).catch((e) => console.error(e));
};

const postEnroute = (rideId: string, coordinate: Coordinate) => {
  setSimulatorStartCoordinate(coordinate);
  void fetchChairPostRideStatus({
    body: { status: "ENROUTE" },
    pathParams: {
      rideId,
    },
  }).catch((e) => console.error(e));
};

const postCarring = (rideId: string) => {
  void fetchChairPostRideStatus({
    body: { status: "CARRYING" },
    pathParams: {
      rideId,
    },
  }).catch((e) => console.error(e));
};

const forcePickup = (pickup_coordinate: Coordinate) =>
  setTimeout(() => {
    currentCoodinatePost(pickup_coordinate);
  }, 60_000);

const forceCarry = (pickup_coordinate: Coordinate, rideId: RideId) =>
  setTimeout(() => {
    (async() => {
      void await currentCoodinatePost(pickup_coordinate);
      void await postCarring(rideId);
    })()
  }, 10_000);

const forceArrive = (pickup_coordinate: Coordinate) =>
  setTimeout(() => {
    currentCoodinatePost(pickup_coordinate);
  }, 60_000);

export const useEmulator = () => {
  const { chair, data, setCoordinate, isAnotherSimulatorBeingUsed } =
    useSimulatorContext();
  const { pickup_coordinate, destination_coordinate, ride_id, status } =
    data ?? {};
  useEffect(() => {
    if (!(pickup_coordinate && destination_coordinate && ride_id)) return;
    let timeoutId: ReturnType<typeof setTimeout>;
    switch (status) {
      case "ENROUTE":
        timeoutId = forcePickup(pickup_coordinate);
        break;
      case "PICKUP":
        timeoutId = forceCarry(pickup_coordinate, ride_id);
        break;
      case "CARRYING":
        timeoutId = forceArrive(destination_coordinate);
        break;
    }
    return () => {
      clearTimeout(timeoutId);
    };
  }, [
    isAnotherSimulatorBeingUsed,
    status,
    destination_coordinate,
    pickup_coordinate,
    ride_id,
  ]);

  useEffect(() => {
    if (isAnotherSimulatorBeingUsed) return;
    if (!(chair && data)) {
      return;
    }

    const timeoutId = setTimeout(() => {
      currentCoodinatePost(chair.coordinate);
      try {
        switch (data.status) {
          case "MATCHING":
            postEnroute(data.ride_id, chair.coordinate);
            break;
          case "PICKUP":
            setCoordinate?.(data.pickup_coordinate);
            postCarring(data.ride_id);
            break;
          case "ENROUTE":
            setCoordinate?.(move(chair.coordinate, data.pickup_coordinate));
            break;
          case "CARRYING":
            setCoordinate?.(
              move(chair.coordinate, data.destination_coordinate),
            );
            break;
          case "ARRIVED":
            setCoordinate?.(data.destination_coordinate);
        }
      } catch (e) {
        // statusの更新タイミングの都合で到着状態を期待しているが必ず取れるとは限らない
      }
    }, 1000);

    return () => {
      clearTimeout(timeoutId);
    };
  }, [chair, data, setCoordinate, isAnotherSimulatorBeingUsed]);
};
