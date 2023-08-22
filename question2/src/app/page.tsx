import Link from "next/link";
import { TrainData } from "./types";

const fetchTrainData = async () => {
  try {
    const response = await fetch("http://localhost:8080/schedules", {
      method: "GET",
    });
    const result: TrainData[] = await response.json();
    return result;
  } catch (error) {
    console.log("error", error);
    throw error;
  }
};

// Example of how to use the fetchTrainData function

const TrainList = async () => {
  const trains = await fetchTrainData();
  return (
    <div className="flex flex-col items-center space-y-4">
      {trains.map((train) => (
        <Link
          key={train.trainNumber}
          href={`/train/${train.trainNumber}`}
          className="border p-4 rounded-md hover:bg-gray-100"
        >
          <h2 className="text-lg font-semibold">{train.trainName}</h2>
          <p>Train Number: {train.trainNumber}</p>
          <p>
            Departure Time:{" "}
            {calculateDepartureTime(train.departureTime, train.delayedBy)}
          </p>
        </Link>
      ))}
    </div>
  );
};

const calculateDepartureTime = (
  departureTime: { Hours: number; Minutes: number; Seconds: number },
  delay: number,
) => {
  const { Hours, Minutes, Seconds } = departureTime;
  const totalDelayInSeconds = delay * 60;

  const departureInSeconds = Hours * 3600 + Minutes * 60 + Seconds;
  const adjustedDepartureTimeInSeconds =
    departureInSeconds + totalDelayInSeconds;

  const adjustedHours = Math.floor(adjustedDepartureTimeInSeconds / 3600) % 24;
  const adjustedMinutes = Math.floor(
    (adjustedDepartureTimeInSeconds % 3600) / 60,
  );
  const adjustedSeconds = adjustedDepartureTimeInSeconds % 60;

  return `${adjustedHours}:${adjustedMinutes}:${adjustedSeconds}`;
};
export default TrainList;
