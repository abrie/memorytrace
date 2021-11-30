export interface Position {
  status: string;
  latitude: number | null;
  longitude: number | null;
}

function ErrorCodeToString(code: number): string {
  switch (code) {
    case 1:
      return 'DENIED';
    case 2:
      return 'UNAVAILABLE';
    case 3:
      return 'TIMEOUT';
  }
  return 'UNKNOWN';
}

export function getPosition(options?: PositionOptions): Promise<Position> {
  return new Promise((resolve, reject) =>
    navigator.geolocation.getCurrentPosition(
      (geo: GeolocationPosition) =>
        resolve({
          status: 'OK',
          latitude: geo.coords.latitude,
          longitude: geo.coords.longitude,
        }),
      (err: GeolocationPositionError) =>
        reject({
          status: ErrorCodeToString(err.code),
          latitude: null,
          longitude: null,
        }),
      options,
    ),
  );
}
