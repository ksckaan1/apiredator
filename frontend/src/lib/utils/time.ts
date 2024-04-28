
export const convertSeconds = (seconds: number): string => {
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const remainingSeconds = seconds % 60;

  if (hours > 0) {
    return `${hours}:${String(minutes).padStart(2, "0")}:${String(remainingSeconds).padStart(2, "0")}`;
  }

  if (minutes > 0) {
    return `${String(minutes).padStart(2, "0")}:${String(remainingSeconds).padStart(2, "0")}`;
  }

  return `${String(remainingSeconds).padStart(2, "0")}`;
};

const months: any = {
  "01": "Jan",
  "02": "Feb",
  "03": "Mar",
  "04": "Apr",
  "05": "May",
  "06": "Jun",
  "07": "Jul",
  "08": "Aug",
  "09": "Sep",
  "10": "Oct",
  "11": "Nov",
  "12": "Dec"
}

export const prettyTime = (time: string): string => {
  if (!/[0-9]{4}/.test(time)) {
    return "";
  }

  if (time === "0001-01-01T00:00:00Z") {
    return "";
  }

  let date = time.split("-");
  let year = date[0];
  let month = date[1];
  let day = date[2].split("T")[0];

  let hours = date[2].split("T")[1].split(":")[0];
  let minutes = date[2].split("T")[1].split(":")[1];
  let seconds = date[2].split("T")[1].split(":")[2].split(".")[0];

  return `${hours}:${minutes}:${seconds} ${months[month]} ${day}, ${year}`;
};