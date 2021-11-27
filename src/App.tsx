import React, { useState } from 'react';
import './App.css';

interface AppProps {}

interface Memory {
  timestamp: number;
  text: string;
}

interface Position {
  longitude: number;
  latitude: number;
}

interface StoredMemory {
  memory: Memory;
  position: Position;
}

function getPosition(options?: PositionOptions): Promise<GeolocationPosition> {
  return new Promise((resolve, reject) =>
    navigator.geolocation.getCurrentPosition(resolve, reject, options),
  );
}

function App({}: AppProps) {
  const [memory, setMemory] = useState<Memory | null>(null);
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

  const storeMemory = () => {
    if (memory == null || position == null) {
      return;
    }
    const str = localStorage.getItem('memories');
    const stored: StoredMemory[] = str ? JSON.parse(str) : [];
    stored.push({ memory, position });
    localStorage.setItem('memories', JSON.stringify(stored));
    setMemory(null);
    setPosition(null);
  };

  return (
    <div className="container w-full border mx-auto">
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
