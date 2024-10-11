import { IconType } from "~/types";

export const PinIcon: IconType<{ color: string }> = function (props) {
  return (
    <svg viewBox="0 0 282 369.24" xmlns="http://www.w3.org/2000/svg" {...props}>
      <path
        d="m281.98 138.82c-1.15-76.2-63.5-138.13-139.71-138.82-78.45-.68-142.27 62.71-142.27 141.01 0 18.6 3.6 36.35 10.14 52.61 2.8 6.94 6.13 13.61 9.94 19.95l12.47 17.55 92.15 129.71c7.97 11.22 24.63 11.22 32.6 0l92.15-129.71 12.47-17.55c3.81-6.34 7.14-13.01 9.94-19.95 6.8-16.9 10.42-35.4 10.12-54.8zm-140.98 44.13c-23.47 0-42.5-19.03-42.5-42.5s19.03-42.5 42.5-42.5 42.5 19.03 42.5 42.5-19.03 42.5-42.5 42.5z"
        fill="currentColor"
      />
    </svg>
  );
};
