import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import type { IconType } from "react-icons";

type ClusterLinkProps = {
  url: string;
  label: string;
  colorClass: string;
  Icon: IconType;
};

export const ClusterLink = ({
  url,
  label,
  colorClass,
  Icon,
}: ClusterLinkProps) => (
  <Tooltip>
    <TooltipTrigger asChild>
      <a
        href={url}
        target="_blank"
        rel="noopener noreferrer"
        className={colorClass}
      >
        <Icon className="w-10 h-10" />
      </a>
    </TooltipTrigger>
    <TooltipContent>{label}</TooltipContent>
  </Tooltip>
);
