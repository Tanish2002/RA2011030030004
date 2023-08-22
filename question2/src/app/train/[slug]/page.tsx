import { TrainData } from "../../types";

const fetchTrainData = async (trainNumber: number) => {
  try {
    const response = await fetch(
      `http://localhost:8080/schedules/${trainNumber}`,
      {
        method: "GET",
      },
    );
    const result: TrainData = await response.json();
    return result;
  } catch (error) {
    console.log("error", error);
    throw error;
  }
};
const TrainDetails = async ({ params }: { params: { slug: string } }) => {
  const trainNumber = params.slug;

  const train = await fetchTrainData(Number(trainNumber));

  if (!train) {
    return <p>Train not found.</p>;
  }

  return (
    <div className="border p-4 rounded-md">
      <h1 className="text-xl font-semibold mb-2">{train.trainName}</h1>
      <p>Train Number: {train.trainNumber}</p>
      <p>
        Departure Time:
        {calculateDepartureTime(train.departureTime, train.delayedBy)}
      </p>
      <p>Seats Available (Sleeper): {train.seatsAvailable.sleeper}</p>
      <p>Seats Available (AC): {train.seatsAvailable.AC}</p>
      <p>Price (Sleeper): {train.price.sleeper}</p>
      <p>Price (AC): {train.price.AC}</p>
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

export default TrainDetails;
