import React from 'react';

interface ToolBarProps {
  onNewMemory: () => void;
}

export function ToolBar({ onNewMemory }: ToolBarProps): JSX.Element {
  return (
    <button
      className="rounded-sm border p-3 bg-green-300"
      onClick={onNewMemory}
    >
      +
    </button>
  );
}
