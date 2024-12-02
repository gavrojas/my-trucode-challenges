// https://thisdavej.com/making-interactive-node-js-console-apps-that-listen-for-keypress-events/
import readline from "readline";

readline.emitKeypressEvents(process.stdin);
process.stdin.setRawMode(true);

export function escapeListener(callback: Function) {
  process.stdin.on("keypress", (_, key) => {
    if (key.name === "escape") {
      callback();
    }
  });
}
