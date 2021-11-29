import React, { useState } from 'react';
import axios from 'axios';

interface AppProps {}

interface Memory {
  timestamp: number;
  text: string;
}

interface Position {
  longitude: number;
  latitude: number;
}

type StoredMemory = Memory & Position;

function getPosition(options?: PositionOptions): Promise<GeolocationPosition> {
  return new Promise((resolve, reject) =>
    navigator.geolocation.getCurrentPosition(resolve, reject, options),
  );
}

function App({}: AppProps) {
  const [memory, setMemory] = useState<Memory | null>(null);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [position, setPosition] = useState<Position | null>(null);

  const newMemory = async () => {
    setMemory({
      timestamp: Date.now(),
      text: '',
    });

    const position = await getPosition();
    setPosition({
      longitude: position.coords.longitude,
      latitude: position.coords.latitude,
    });
  };

  const cancelMemory = () => {
    setMemory(null);
    setPosition(null);
  };

  const storeMemory = async () => {
    if (memory == null || position == null) {
      return;
    }
    const storedMemory: StoredMemory = {
      ...memory,
      ...position,
    };
    try {
      await axios.post('/api/_memory', storedMemory);
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
    setMemory(null);
    setPosition(null);
  };

  return (
    <div className="container w-full m-2 mx-auto">
      <Title />
      <ErrorBanner message={errorMessage} />
      <ToolBar onNewMemory={newMemory} />
      <MemoryEditor
        memory={memory}
        position={position}
        setMemory={setMemory}
        storeMemory={storeMemory}
        cancelMemory={cancelMemory}
      />
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
  position: Position | null;
  memory: Memory | null;
  setMemory: (memory: Memory) => void;
  storeMemory: () => void;
  cancelMemory: () => void;
}

function MemoryEditor({
  memory,
  position,
  setMemory,
  storeMemory,
  cancelMemory,
}: MemoryEditorProps): JSX.Element {
  if (memory === null) {
    return <></>;
  }
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
        <div className="px-5">{position ? 'ok' : 'acquiring...'}</div>
        <button className="bg-green-400 px-5" onClick={() => storeMemory()}>
          store
        </button>
        <button className="px-5 bg-yellow-400" onClick={() => cancelMemory()}>
          cancel
        </button>
      </div>
    </div>
  );
}

export default App;
