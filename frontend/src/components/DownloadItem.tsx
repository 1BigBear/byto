import { useState } from 'react';
import { Play, Pause, Square, Trash2, ChevronDown, ChevronUp, CheckCircle, Loader2, AlertCircle, FolderOpen } from 'lucide-react';
import { Button } from './ui/button';
import { Progress } from './ui/progress';

interface DownloadVideo {
  id: string;
  url: string;
  fileName: string;
  progress: number;
  fileSize: string;
  status: 'pending' | 'downloading' | 'paused' | 'completed' | 'error';
  logs: string[];
}

interface DownloadItemProps {
  download: DownloadVideo;
  onAction: (id: string, action: 'start' | 'pause' | 'resume' | 'delete') => void;
  onRemove: (id: string) => void;
  onShowInFolder?: (id: string) => void;
}

export function DownloadItem({ download, onAction, onRemove, onShowInFolder }: DownloadItemProps) {
  const [showLogs, setShowLogs] = useState(false);

  const getStatusIcon = () => {
    switch (download.status) {
      case 'downloading':
        return <Loader2 className="size-4 text-blue-500 animate-spin" />;
      case 'completed':
        return <CheckCircle className="size-4 text-green-500" />;
      case 'error':
        return <AlertCircle className="size-4 text-red-500" />;
      case 'paused':
        return <Pause className="size-4 text-orange-500" />;
      default:
        return <div className="size-4 rounded-full border-2 border-gray-600" />;
    }
  };

  const getStatusColor = () => {
    switch (download.status) {
      case 'downloading':
        return 'bg-blue-500';
      case 'completed':
        return 'bg-green-500';
      case 'error':
        return 'bg-red-500';
      case 'paused':
        return 'bg-orange-500';
      default:
        return 'bg-gray-600';
    }
  };

  return (
    <div className="bg-[#141414] rounded-lg border border-[#262626] overflow-hidden">
      <div className="p-4">
        <div className="flex items-start gap-3">
          <div className="mt-1">{getStatusIcon()}</div>
          
          <div className="flex-1 min-w-0 space-y-3">
            {/* File Name and URL */}
            <div>
              <h3 className="text-gray-100 truncate">{download.fileName}</h3>
              <p className="text-gray-500 text-sm truncate">{download.url}</p>
            </div>

            {/* Progress Bar */}
            <div className="space-y-1">
              <div className="flex justify-between text-sm">
                <span className="text-gray-400">{download.status === 'completed' ? 'Completed' : download.status === 'downloading' ? 'Downloading...' : download.status === 'paused' ? 'Paused' : 'Pending'}</span>
                <div className="flex items-center gap-3">
                  <span className="text-gray-400">{download.fileSize}</span>
                  <span className="text-xs text-gray-300">{download.progress}%</span>
                </div>
              </div>
              <Progress value={download.progress} className="h-2" />
            </div>

            {/* Logs Toggle */}
            {download.logs.length > 0 && (
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setShowLogs(!showLogs)}
                className="text-gray-400 px-0 hover:text-gray-300"
              >
                {showLogs ? <ChevronUp className="size-4" /> : <ChevronDown className="size-4" />}
                {showLogs ? 'Hide Logs' : 'Show Logs'}
              </Button>
            )}
          </div>

          {/* Action Buttons */}
          <div className="flex gap-1">
            {download.status === 'pending' || download.status === 'paused' ? (
              <Button
                variant="outline"
                size="icon"
                onClick={() => onAction(download.id, 'start')}
                title="Start Download"
                className="border-[#262626] hover:bg-[#1f1f1f]"
              >
                <Play className="size-4" />
              </Button>
            ) : download.status === 'downloading' ? (
              <Button
                variant="outline"
                size="icon"
                onClick={() => onAction(download.id, 'pause')}
                title="Pause Download"
                className="border-[#262626] hover:bg-[#1f1f1f]"
              >
                <Pause className="size-4" />
              </Button>
            ) : null}
            
            <Button
              variant="outline"
              size="icon"
              onClick={() => onRemove(download.id)}
              title="Remove from List"
              className="border-[#262626] hover:bg-[#1f1f1f] hover:text-red-500"
            >
              <Trash2 className="size-4" />
            </Button>

            {onShowInFolder && download.status === 'completed' && (
              <Button
                variant="outline"
                size="icon"
                onClick={() => onShowInFolder(download.id)}
                title="Show in Folder"
                className="border-[#262626] hover:bg-[#1f1f1f] hover:text-blue-500"
              >
                <FolderOpen className="size-4" />
              </Button>
            )}
          </div>
        </div>
      </div>

      {/* Logs Section */}
      {showLogs && download.logs.length > 0 && (
        <div className="border-t border-[#262626] bg-[#0a0a0a] p-4">
          <div className="bg-black rounded text-green-400 p-3 font-mono text-xs space-y-1 max-h-48 overflow-y-auto">
            {download.logs.map((log, index) => (
              <div key={index}>{log}</div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}