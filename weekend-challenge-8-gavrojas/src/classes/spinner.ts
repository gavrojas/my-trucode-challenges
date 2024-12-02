// spinner.ts
import { IntervalOrNull, EasterEgg } from "../../types";
import { escapeListener } from "../utils/keypress";
import chalk from "chalk";
import { Loader } from "./loader";

export class Spinner implements EasterEgg {
  interval: IntervalOrNull;
  onComplete: () => void;
  loader: Loader;
  currentFrame: number;
  frames: string[];
  colors: any[];
  currentColor: number;

  constructor(onComplete: () => void) {
    this.interval = null;
    this.onComplete = onComplete;
    this.loader = new Loader(() => this.animate());
    this.currentFrame = 0;
    this.frames = ["-", "\\", "|", "/"];
    this.colors = [chalk.red, chalk.magenta, chalk.blue];
    this.currentColor = 0;
  }

  init() {
    this.load();
  }

  draw() {
    console.clear();
    const color = this.colors[this.currentColor];
    console.log(color(this.frames[this.currentFrame]));

    this.currentFrame++;
    if (this.currentFrame >= this.frames.length) {
      this.currentFrame = 0;
    }

    this.currentColor++;
    if (this.currentColor >= this.colors.length) {
      this.currentColor = 0;
    }

  }

  animate() {
    this.interval = setInterval(this.draw.bind(this), 100);
    escapeListener(() => this.end());
  }

  end() {
    if (this.interval) {
      clearInterval(this.interval);
      this.interval = null;
    }
    console.clear();
    this.onComplete();
    process.stdin.setRawMode(false);
    process.exit(0);
  }

  load() {
    this.loader.load();
  }
}
