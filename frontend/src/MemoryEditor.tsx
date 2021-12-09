import React, { useState, useEffect, useRef } from 'react';

import { getPosition } from './geolocation';
import type { Position } from './geolocation';
import type { Memory } from './models';

interface MemoryEditorProps {
  storeMemory: (memory: Memory) => void;
  cancelMemory: () => void;
}

type Link = string;
type Chain = Link[];

interface ChainViewProps {
  chain: Chain;
}

function ChainView(props: ChainViewProps): JSX.Element {
  return (
    <>
      {props.chain.map((link: Link, idx: number) => (
        <div key={idx}>{link}</div>
      ))}
    </>
  );
}

export function MemoryEditor(props: MemoryEditorProps): JSX.Element {
  const inputRef = useRef<HTMLTextAreaElement | null>(null);
  const [text, setText] = useState<string>('');
  const [chain, setChain] = useState<Chain>([]);
  const [timestamp] = useState<number>(Date.now());
  const [position, setPosition] = useState<Position>({
    status: 'PENDING',
    latitude: null,
    longitude: null,
  });

  function addLink(text: string) {
    setChain([...chain, text]);
    setText('');
    if (inputRef.current !== null) {
      inputRef.current.focus();
    }
  }

  const buildMemory = () => {
    // Copy the last link text from the edit, if present.
    const finalChain = text === '' ? chain : [...chain, text];
    return {
      chain: finalChain,
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
      <ChainView chain={chain} />
      <div className="flex flex-row">
        <textarea
          className="p-1 w-full border-2"
          placeholder="...???"
          onChange={(e: React.FormEvent<HTMLTextAreaElement>) => {
            setText(e.currentTarget.value);
          }}
          value={text}
          ref={inputRef}
          autoFocus
        ></textarea>
        <button onClick={() => addLink(text)}>next</button>
      </div>
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
