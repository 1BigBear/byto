import { Star, Github, Heart } from 'lucide-react';
import { Button } from './ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from './ui/dialog';
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime';

interface SupportPanelProps {
  onClose: () => void;
}

export function SupportPanel({ onClose }: SupportPanelProps) {
  const openLink = (url: string) => {
    BrowserOpenURL(url);
  };

  return (
    <Dialog open={true} onOpenChange={onClose}>
      <DialogContent className="max-w-md bg-[#141414] border border-[#262626] text-gray-100">
        <DialogHeader>
          <DialogTitle className="text-gray-100">Support Byto</DialogTitle>
          <DialogDescription className="text-gray-400">
            If you find Byto useful, consider supporting the project
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-3 py-4">
          <Button
            variant="outline"
            className="w-full justify-start gap-2 border-[#262626] hover:bg-[#1f1f1f]"
            onClick={() => openLink('https://github.com/OmarNaru1110/byto')}
          >
            <Star className="size-4" />
            Star byto on GitHub
          </Button>
          <Button
            variant="outline"
            className="w-full justify-start gap-2 border-[#262626] hover:bg-[#1f1f1f]"
            onClick={() => openLink('https://github.com/OmarNaru1110')}
          >
            <Github className="size-4" />
            Follow me on GitHub
          </Button>
          <Button
            variant="outline"
            className="w-full justify-start gap-2 border-[#262626] hover:bg-[#1f1f1f]"
            onClick={() => openLink('https://ko-fi.com/omarnaru')}
          >
            <Heart className="size-4" />
            Support me on Ko-fi
          </Button>
        </div>

        <div className="flex justify-end pt-4 border-t border-[#262626]">
          <Button onClick={onClose} className="bg-blue-600 hover:bg-blue-700">
            Close
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
}
