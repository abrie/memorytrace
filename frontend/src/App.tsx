import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface AppProps {}

interface Memory {
  timestamp: number;
  text: string;
  longitude: number | null;
  latitude: number | null;
}

function getPosition(options?: PositionOptions): Promise<GeolocationPosition> {
  return new Promise((resolve, reject) =>
    navigator.geolocation.getCurrentPosition(resolve, reject, options),
  );
}

function App({}: AppProps) {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [showEditor, setShowEditor] = useState<boolean>(false);

  const storeMemory = async (memory: Memory) => {
    try {
      await axios.post('/api/memory', memory);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        setErrorMessage(error.message);
        return;
      } else {
        setErrorMessage('Client Error: Unable to create request.');
        return;
      }
    }
    setErrorMessage(null);
    setShowEditor(false);
  };

  return (
    <div className="container w-full m-2 mx-auto">
      <Title />
      <ErrorBanner message={errorMessage} />
      <ToolBar onNewMemory={() => setShowEditor(true)} />
      {showEditor && (
        <MemoryEditor
          storeMemory={storeMemory}
          cancelMemory={() => setShowEditor(false)}
        />
      )}
    </div>
  );
}

function Title(): JSX.Element {
  return (
    <header>
      <div className="font-mono text-center">memorytrace</div>
    </header>
  );
}

interface ToolBarProps {
  onNewMemory: () => void;
}

function ToolBar({ onNewMemory }: ToolBarProps): JSX.Element {
  return (
    <button
      className="rounded-sm border p-3 bg-green-300"
      onClick={onNewMemory}
    >
      +
    </button>
  );
}

interface ErrorBannerProps {
  message: string | null;
}

function ErrorBanner({ message }: ErrorBannerProps): JSX.Element {
  if (message === null) {
    return <></>;
  } else {
    return <div className="p-1 bg-red-500 text-white font-bold">{message}</div>;
  }
}
interface MemoryEditorProps {
  storeMemory: (memory: Memory) => void;
  cancelMemory: () => void;
}

function MemoryEditor(props: MemoryEditorProps): JSX.Element {
  const newMemory: Memory = {
    text: '',
    timestamp: Date.now(),
    latitude: null,
    longitude: null,
  };

  const [memory, setMemory] = useState<Memory>(newMemory);
  const [geoAcquired, setGeoAcquired] = useState<boolean>(false);

  useEffect(() => {
    async function aquire() {
      const geo = await getPosition();
      setMemory({
        ...memory,
        longitude: geo.coords.longitude,
        latitude: geo.coords.latitude,
      });
      setGeoAcquired(true);
    }
    aquire();
  });

  return (
    <div>
      <textarea
        className="p-1 w-full border-2"
        placeholder="...???"
        onChange={(e: React.FormEvent<HTMLTextAreaElement>) => {
          setMemory({ ...memory, text: e.currentTarget.value });
        }}
        value={memory.text}
        autoFocus
      ></textarea>
      <div className="flex flex-row">
        <div className="px-5">{geoAcquired ? 'ok' : 'acquiring...'}</div>
        <button
          className="bg-green-400 px-5"
          onClick={() => props.storeMemory(memory)}
        >
          store
        </button>
        <button
          className="px-5 bg-yellow-400"
          onClick={() => props.cancelMemory()}
        >
          cancel
        </button>
      </div>
    </div>
  );
}

export default App;
