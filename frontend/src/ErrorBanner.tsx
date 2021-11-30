import React from 'react';

interface ErrorBannerProps {
  message: string | null;
}

export function ErrorBanner({ message }: ErrorBannerProps): JSX.Element {
  if (message === null) {
    return <></>;
  } else {
    return <div className="p-1 bg-red-500 text-white font-bold">{message}</div>;
  }
}
