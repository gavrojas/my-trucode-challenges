import chalk from "chalk";
import { EasterEggOrUndefined, EASTER_EGGS } from "../types";
import { Loader } from "./classes/loader";
import { Shark } from "./classes/shark";
import { Clock } from "./classes/clock";
import { Bounce } from "./classes/bounce";
import { Marquee } from "./classes/marquee";
import { Spinner } from "./classes/spinner";


const easterEggsDict: { [key: string]: EasterEggOrUndefined } = {
  loader: new Loader(() => console.log("Loader done!")),
  shark: new Shark(() => console.log("Shark ended. Thanks for using my app!")),
  clock: new Clock (() => console.log("Clock ended. Thanks for using my app!")),
  bounce: new Bounce (() => console.log("Bounce ended. Thanks for using my app!")),
  marquee: new Marquee (() => console.log("Marquee ended. Thanks for using my app!")),
  spinner: new Spinner (() => console.log("Spinner ended. Thanks for using my app!")),
};

(function () {
  if (process.argv[2] === "--egg") {
    const eggFlag: string = process.argv[3];
    const animationInstance = easterEggsDict[eggFlag];

    if (animationInstance) {
      animationInstance.init();
    } else {
      console.log(chalk.red("Are you sure that's an egg? O_o"));
    }
    return;
  }
  console.log(chalk.green("Hello world"));
})();
