import { IntervalOrNull, EasterEgg  } from "../../types";
import { escapeListener } from "../utils/keypress";
import { Loader } from "./loader";

export class Shark implements EasterEgg {
  interval: IntervalOrNull;
  currentFrame: number;
  frames: string[];
  onComplete: () => void;
  loader: Loader;

  constructor(onComplete: () => void) {
    this.interval = null;
    this.currentFrame = 0;
    this.frames = [
      "|\\_________________________________",
      "_|\\________________________________",
      "__|\\_______________________________",
      "___|\\______________________________",
      "____|\\_____________________________",
      "_____|\\____________________________",
      "______|\\___________________________",
      "_______|\\__________________________",
      "________|\\_________________________",
      "_________|\\________________________",
      "__________|\\_______________________",
      "___________|\\______________________",
      "____________|\\_____________________",
      "_____________|\\____________________",
      "______________|\\___________________",
      "_______________|\\__________________",
      "________________|\\_________________",
      "_________________|\\________________",
      "__________________|\\_______________",
      "___________________|\\______________",
      "____________________|\\_____________",
      "_____________________|\\____________",
      "______________________|\\___________",
      "_______________________|\\__________",
      "________________________|\\_________",
      "_________________________|\\________",
      "__________________________|\\_______",
      "___________________________|\\______",
      "____________________________|\\_____",
      "_____________________________|\\____",
      "______________________________|\\___",
      "_______________________________|\\__",
      "________________________________|\\_",
      "_________________________________|\\",
    ];
    this.onComplete = onComplete;
    this.loader = new Loader(() => this.animate());
  }

  draw() {
    console.clear();
    console.log(this.frames[this.currentFrame]);
    this.currentFrame++;
    if (this.currentFrame === this.frames.length) this.currentFrame = 0;
  }

  animate() {
    this.interval = setInterval(this.draw.bind(this), 300);
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