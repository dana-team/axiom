import { useEffect, useState } from "react";
import ReactJson from "react-json-view";

interface Props {
  value: string;
}

const ComparisonDisplay = ({ value }: Props) => {
  const [jsonTheme, setJsonTheme] = useState<"rjv-default" | "brewer">(
    "rjv-default"
  );

  useEffect(() => {
    const isDark = document.documentElement.classList.contains("dark");
    setJsonTheme(isDark ? "brewer" : "rjv-default");
  }, []);

  return (
    <div className="text-sm h-131 overflow-auto">
      <ReactJson
        src={JSON.parse(value)}
        name={false}
        enableClipboard={false}
        displayDataTypes={false}
        theme={jsonTheme}
      />
    </div>
  );
};

export default ComparisonDisplay;
