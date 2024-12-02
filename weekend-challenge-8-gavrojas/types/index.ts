export enum EASTER_EGGS {
  shark,
  clock,
  spinner,
  bounce,
  marquee,
}

export interface EasterEgg {
  draw(): void;
  animate(): void;
  init(): void;
  end(): void;
}

export type EasterEggOrUndefined = EasterEgg | undefined;
export type IntervalOrNull = NodeJS.Timeout | null;
