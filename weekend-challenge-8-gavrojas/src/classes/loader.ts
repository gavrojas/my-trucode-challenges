import { IntervalOrNull, EasterEgg  } from "../../types";

export class Loader implements EasterEgg {
  interval: IntervalOrNull;
  currentFrame: number;
  frames: string[];
  onComplete: () => void;

  constructor(onComplete: () => void) {
    this.interval = null;
    this.currentFrame = 0;
    this.frames = [".", "..", "..."];
    this.onComplete = onComplete;
  }

  init() {
    this.load();
  }

  draw() {
    console.clear();
    console.log(this.frames[this.currentFrame]);
    this.currentFrame++;
    if (this.currentFrame === this.frames.length) this.currentFrame = 0;
  }

  animate() {
    this.interval = setInterval(this.draw.bind(this), 300);
  }

  end() {
    if (this.interval) {
      clearInterval(this.interval);
      this.interval = null;
    }
    console.clear();
    this.onComplete();
  }

  load() {
    this.animate();
    setTimeout(() => {
      this.end();
    }, 3000);
  }
}
