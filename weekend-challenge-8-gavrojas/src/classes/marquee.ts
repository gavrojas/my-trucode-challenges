// marquee.ts
import { IntervalOrNull, EasterEgg } from "../../types";
import chalk from "chalk";
import { escapeListener } from "../utils/keypress";
import { Loader } from "./loader";

export class Marquee implements EasterEgg {
  interval: IntervalOrNull;
  text: string;
  displayText: string;
  onComplete: () => void;
  loader: Loader;
  currentIndex: number;
  colorText: boolean;

  constructor(onComplete: () => void) {
    this.interval = null;
    this.text = "Trucode 2! :D";
    this.displayText = " ".repeat(36) + this.text;
    this.onComplete = onComplete;
    this.loader = new Loader(() => this.animate());
    this.currentIndex = 0;
    this.colorText = true;
  }

  init() {
    this.load();
  }

  draw() {
    console.clear();
    const visibleText = this.displayText.slice(this.currentIndex, this.currentIndex + 36);
    const color = this.colorText ? chalk.magenta : chalk.blue;
    console.log(color(visibleText));
    this.colorText = !this.colorText;
    this.currentIndex++;
    if (this.currentIndex >= this.displayText.length) {
      this.currentIndex = 0;
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
