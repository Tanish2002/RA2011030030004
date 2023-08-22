
export interface TrainData {
  trainName: string;
  trainNumber: string;
  departureTime: {
    Hours: number;
    Minutes: number;
    Seconds: number;
  };
  seatsAvailable: {
    sleeper: number;
    AC: number;
  };
  price: {
    sleeper: number;
    AC: number;
  };
  delayedBy: number;
}
