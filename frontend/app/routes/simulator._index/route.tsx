import type { MetaFunction } from "@remix-run/react";
import { useEffect, useRef } from "react";
import { fetchChairPostActivity } from "~/apiClient/apiComponents";
import { useEmulator } from "~/components/hooks/use-emulator";
import { SimulatorChairDisplay } from "~/components/modules/simulator-display/simulator-chair-display";
import { SimulatorConfigDisplay } from "~/components/modules/simulator-display/simulator-config-display";
import { SmartPhone } from "~/components/primitives/smartphone/smartphone";
import { useSimulatorContext } from "~/contexts/simulator-context";

export const meta: MetaFunction = () => {
  return [
    { title: "Simulator | ISURIDE" },
    { name: "description", content: "isucon14" },
  ];
};

export default function Index() {
  const { targetChair: chair } = useSimulatorContext();
  const ref = useRef<HTMLIFrameElement>(null);

  useEmulator(chair);

  useEffect(() => {
    try {
      void fetchChairPostActivity({ body: { is_active: true } });
    } catch (error) {
      console.error(error);
    }
  }, []);

  return (
    <main className="h-screen flex justify-center items-center space-x-8 lg:space-x-16">
      <SmartPhone>
        <iframe
          title="ISURIDE Client App"
          src="/client"
          className="w-full h-full"
          ref={ref}
        />
      </SmartPhone>
      <div className="space-y-4 min-w-[320px] lg:w-[400px]">
        <h1 className="text-lg font-semibold mb-4">Chair Simulator</h1>
        <SimulatorChairDisplay />
        <SimulatorConfigDisplay simulatorRef={ref} />
      </div>
    </main>
  );
}
