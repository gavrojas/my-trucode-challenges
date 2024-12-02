// clock.ts
import { IntervalOrNull, EasterEgg } from "../../types";
import { escapeListener } from "../utils/keypress";
import chalk from "chalk";
import { Loader } from "./loader";

export class Clock implements EasterEgg {
  interval: IntervalOrNull;
  onComplete: () => void;
  blinkingColon: boolean;
  loader: Loader;

  constructor(onComplete: () => void) {
    this.interval = null;
    this.onComplete = onComplete;
    this.blinkingColon = true;
    this.loader = new Loader(() => this.startClock());
  }

  init() {
    this.load();
  }

  draw() {
    const now = new Date();
    const hours = String(now.getHours()).padStart(2, "0");
    const minutes = String(now.getMinutes()).padStart(2, "0");
    const seconds = String(now.getSeconds()).padStart(2, "0");
    const separator = this.blinkingColon ? ":" : " ";
    this.blinkingColon = !this.blinkingColon;
    console.clear();
    console.log(chalk.green(`${hours}${separator}${minutes}${separator}${seconds}`));
  }

  animate() {
    this.interval = setInterval(this.draw.bind(this), 500);
    escapeListener(() => this.end());
  }

  end() {
    if (this.interval) {
      clearInterval(this.interval);
      this.interval = null;
    }
    console.clear();
    this.onComplete();
    // asegurarse que termina porque la terminal se quedaba pegada
    process.stdin.setRawMode(false);
    process.exit(0);
  }

  load() {
    this.loader.load();
  }

  startClock() {
    this.animate();
  }
}
