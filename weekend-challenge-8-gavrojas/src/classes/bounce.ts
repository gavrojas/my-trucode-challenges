// bounce.ts
import { IntervalOrNull, EasterEgg } from "../../types";
import { escapeListener } from "../utils/keypress";
import chalk from "chalk";
import { Loader } from "./loader";

export class Bounce implements EasterEgg {
  interval: IntervalOrNull;
  currentFrame: number;
  frames: string[];
  onComplete: () => void;
  loader: Loader;

  constructor(onComplete: () => void) {
    this.interval = null;
    this.currentFrame = 0;
    this.frames = this.generateFrames();
    this.onComplete = onComplete;
    this.loader = new Loader(() => this.animate());
  }

  generateFrames(): string[] {
    const frames: string[] = [];
    const arrowRight = "=>";
    const reverseLeft = "<=";
    for (let i = 0; i < 36; i++) {
      frames.push(" ".repeat(i) + arrowRight);
    }
    for (let i = 35; i >= 0; i--) {
      frames.push(" ".repeat(i) + reverseLeft);
    }
    return frames;
  }

  draw() {
    console.clear();
    console.log(chalk.blue(this.frames[this.currentFrame]));
    this.currentFrame++;
    if (this.currentFrame === this.frames.length) this.currentFrame = 0;
  }

  animate() {
    this.interval = setInterval(this.draw.bind(this), 50);
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

  init() {
    this.load();
  }

  load() {
    this.loader.load();
  }
}
