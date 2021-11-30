import React, { useState, useEffect } from 'react';

import { getPosition } from './geolocation';
import type { Position } from './geolocation';
import type { Memory } from './models';

interface MemoryEditorProps {
  storeMemory: (memory: Memory) => void;
  cancelMemory: () => void;
}

export function MemoryEditor(props: MemoryEditorProps): JSX.Element {
  const [text, setText] = useState<string>('');
  const [timestamp] = useState<number>(Date.now());
  const [position, setPosition] = useState<Position>({
    status: 'PENDING',
    latitude: null,
    longitude: null,
  });

  const buildMemory = () => {
    return {
      text,
      timestamp,
      geolocationStatus: position.status,
      latitude: position.latitude,
      longitude: position.longitude,
    };
  };
  useEffect(() => {
    async function aquire() {
      const p = await getPosition();
      setPosition(p);
    }
    aquire();
  }, []);

  return (
    <div>
      <textarea
        className="p-1 w-full border-2"
        placeholder="...???"
        onChange={(e: React.FormEvent<HTMLTextAreaElement>) => {
          setText(e.currentTarget.value);
        }}
        value={text}
        autoFocus
      ></textarea>
      <div className="flex flex-row">
        <button
          className="bg-green-400 px-5"
          onClick={() => props.storeMemory(buildMemory())}
        >
          store
        </button>
        <button
          className="px-5 bg-yellow-400"
          onClick={() => props.cancelMemory()}
        >
          cancel
        </button>
        <div className="px-5">{position.status}</div>
      </div>
    </div>
  );
}
