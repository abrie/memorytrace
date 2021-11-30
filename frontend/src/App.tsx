import React, { useState } from 'react';
import { MemoryEditor } from './MemoryEditor';
import { TitleBar } from './TitleBar';
import { ErrorBanner } from './ErrorBanner';
import { ToolBar } from './ToolBar';
import type { Memory } from './models';
import axios from 'axios';

interface AppProps {}

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
      <TitleBar />
      <ErrorBanner message={errorMessage} />
      <ToolBar
        canAddNewMemory={!showEditor}
        onNewMemory={() => setShowEditor(true)}
      />
      {showEditor && (
        <MemoryEditor
          storeMemory={storeMemory}
          cancelMemory={() => setShowEditor(false)}
        />
      )}
    </div>
  );
}

export default App;
