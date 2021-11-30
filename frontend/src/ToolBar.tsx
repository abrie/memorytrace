import React from 'react';

interface ToolBarProps {
  canAddNewMemory: boolean;
  onNewMemory: () => void;
}

export function ToolBar({
  canAddNewMemory,
  onNewMemory,
}: ToolBarProps): JSX.Element {
  return (
    <>
      {canAddNewMemory && (
        <button
          className="rounded-sm border p-3 bg-green-300"
          onClick={onNewMemory}
        >
          +
        </button>
      )}
    </>
  );
}
