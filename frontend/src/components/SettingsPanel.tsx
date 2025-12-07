import { useState } from 'react';
import { X, FolderOpen } from 'lucide-react';
import { Button } from './ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from './ui/dialog';
import { SelectDownloadFolder } from '../../wailsjs/go/main/App';

interface SettingsPanelProps {
  downloadPath: string;
  quality: string;
  parallelDownloads: string;
  onClose: () => void;
  onSave: (settings: { downloadPath: string; quality: string; parallelDownloads: string }) => void;
}

export function SettingsPanel({
  downloadPath: initialDownloadPath,
  quality: initialQuality,
  parallelDownloads: initialParallelDownloads,
  onClose,
  onSave
}: SettingsPanelProps) {
  // Local state - only applied when Save is clicked
  const [localDownloadPath, setLocalDownloadPath] = useState(initialDownloadPath);
  const [localQuality, setLocalQuality] = useState(initialQuality);
  const [localParallelDownloads, setLocalParallelDownloads] = useState(initialParallelDownloads);

  const handleSelectFolder = async () => {
    try {
      const path = await SelectDownloadFolder();
      if (path) {
        setLocalDownloadPath(path);
      }
    } catch (error) {
      console.error('Error selecting folder:', error);
    }
  };

  const handleSave = () => {
    onSave({
      downloadPath: localDownloadPath,
      quality: localQuality,
      parallelDownloads: localParallelDownloads
    });
  };

  return (
    <Dialog open={true} onOpenChange={onClose}>
      <DialogContent className="max-w-4xl bg-[#141414] border border-[#262626] text-gray-100">
        <DialogHeader>
          <DialogTitle className="text-gray-100">Settings</DialogTitle>
          <DialogDescription className="text-gray-400">
            Configure your download preferences
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4 py-4">
          <div>
            <label className="text-gray-300 text-sm">Default Download Path</label>
            <p className="text-gray-500 text-xs mb-2">Where files will be saved</p>
            <div className="flex gap-2">
              <input
                type="text"
                value={localDownloadPath}
                onChange={(e) => setLocalDownloadPath(e.target.value)}
                className="flex-1 px-3 py-2 bg-[#1f1f1f] border border-[#262626] rounded text-sm text-gray-100"
              />
              <Button 
                size="sm" 
                variant="outline" 
                className="border-[#262626] hover:bg-[#1f1f1f]"
                onClick={handleSelectFolder}
              >
                <FolderOpen className="size-4" />
              </Button>
            </div>
          </div>

          <div>
            <label className="text-gray-300 text-sm">Parallel Downloads</label>
            <p className="text-gray-500 text-xs mb-2">Number of simultaneous downloads</p>
            <select
              value={localParallelDownloads}
              onChange={(e) => setLocalParallelDownloads(e.target.value)}
              className="w-full px-3 py-2 bg-[#1f1f1f] border border-[#262626] rounded text-sm text-gray-100"
            >
              <option value="1">1 (Sequential)</option>
              <option value="2">2</option>
              <option value="3">3</option>
              <option value="4">4</option>
              <option value="5">5</option>
              <option value="10">10</option>
            </select>
          </div>

          <div>
            <label className="text-gray-300 text-sm">Default Video Quality</label>
            <p className="text-gray-500 text-xs mb-2">Preferred quality for downloads</p>
            <select
              value={localQuality}
              onChange={(e) => setLocalQuality(e.target.value)}
              className="w-full px-3 py-2 bg-[#1f1f1f] border border-[#262626] rounded text-sm text-gray-100"
            >
              <option value="360p">360p</option>
              <option value="480p">480p</option>
              <option value="720p">720p (HD)</option>
              <option value="1080p">1080p (Full HD)</option>
              <option value="1440p">1440p (2K)</option>
              <option value="2160p">2160p (4K)</option>
              <option value="best">Best Available</option>
            </select>
          </div>
        </div>

        <div className="flex justify-end gap-2 pt-4 border-t border-[#262626]">
          <Button variant="outline" onClick={onClose} className="border-[#262626] hover:bg-[#1f1f1f]">
            Cancel
          </Button>
          <Button onClick={handleSave} className="bg-blue-600 hover:bg-blue-700">
            Save Changes
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
}