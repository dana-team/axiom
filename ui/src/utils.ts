export const capitalize = (s: string): string => s ? s[0].toUpperCase() + s.slice(1) : s;

export const buildConsoleUrl = (name: string): string =>
  "https://console-openshift-console.apps." + name;
