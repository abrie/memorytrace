import React, { useState } from 'react';
import './App.css';
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

  const storeMemory = async () => {
    if (memory == null || position == null) {
      return;
    }
    const storedMemory: StoredMemory = {
      ...memory,
      ...position,
    };
    try {
      await axios.post('/api/memory', storedMemory);
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
      {errorMessage && (
        <div className="bg-red-400 text-white font-bold">{errorMessage}</div>
      )}
      <button className="rounded-sm border" onClick={newMemory}>
        +
      </button>
      {memory && MemoryEditor({ memory, position, setMemory, storeMemory })}
    </div>
  );
}

interface MemoryEditorProps {
  position: Position | null;
  memory: Memory;
  setMemory: (memory: Memory) => void;
  storeMemory: () => void;
}

function MemoryEditor({
  memory,
  position,
  setMemory,
  storeMemory,
}: MemoryEditorProps): JSX.Element {
  return (
    <div>
      <div className="flex flex-row">
        <div>timestamp: {memory.timestamp}</div>
        <div className="mx-1">.</div>
        <div>{position ? 'ok' : 'capturing...'}</div>
        <button onClick={() => storeMemory()}>store</button>
      </div>
      <textarea
        className="p-1 w-full border-2"
        placeholder="...???"
        onChange={(e: React.FormEvent<HTMLTextAreaElement>) => {
          setMemory({ ...memory, text: e.currentTarget.value });
        }}
        value={memory.text}
        autoFocus
      ></textarea>
    </div>
  );
}

export default App;
