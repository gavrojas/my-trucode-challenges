[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/NigAobF1)
# Easter Eggs

---

In the context of software and media, an "Easter egg" is a hidden feature, message, or inside joke that is embedded by the creators for users to discover. Easter eggs are often found in video games, movies, software applications, and websites. They are meant to be a fun and sometimes surprising reward for users who explore the product thoroughly.

Easter eggs can take various forms, such as hidden messages, secret levels in games, or references to popular culture. They are typically not essential to the main functionality of the software or media but add an element of entertainment for those who come across them. The term "Easter egg" is a reference to the Easter egg hunts that are organized during the Easter holiday, where participants search for hidden eggs.

Discovering an Easter egg can be a delightful experience for users, fostering a sense of community and shared enjoyment among those who find and share these hidden gems.

## How it works?

By running `npx tsx app.ts`

> 💡 gif images show the use of ts-node, ignore that and use tsx instead

The program will work as expected and log `Hello World`

But, if we add the flag `--egg` we can discover a brand new world.
If nothing is passed after the --egg flag, then the app will show us:
`"Are you sure that's an egg? O_o"`

![npx tsx app.ts](./captures/Kapture%202023-11-24%20at%2003.23.27.gif)

Otherwise, if we type a valid easter egg name, we will have the next behaviors:

- **Spinner**

  ![npx tsx app.ts --egg spinner](./captures/Kapture%202023-11-24%20at%2003.56.17.gif)
  
- **Shark**

  ![npx tsx app.ts --egg shark](./captures/Kapture%202023-11-24%20at%2003.25.54.gif)
- **Clock**
  
  ![npx tsx app.ts --egg clock](./captures/Kapture%202023-11-24%20at%2003.29.06.gif)
- **Bounce**

  ![npx tsx app.ts --egg bounce](./captures/Kapture%202023-11-24%20at%2003.31.15.gif)
- **Marquee**

  ![npx tsx app.ts --egg clock](./captures/Kapture%202023-11-24%20at%2003.53.56.gif)

## Directives

1. The order presented above is the recommended order of difficulty.
2. Before commencing development, ensure that you install the required packages by executing the command `npm install`.
3. Refer to the accompanying gifs in this document for guidance on running your code.
4. All code must be placed within the `src` directory.
5. In the event of an invalid easter egg, display the message `"Are you sure that's an egg? O_o"`.
6. If you choose to incorporate colors, you have the option of utilizing the `chalk` library, which is already included in this package. Note that this is not mandatory.
7. All animations should run infinitely unless the `escape` key is pressed. Refer to the `escapeListener` function in the `utils` folder for implementation details.
8. Within the `types` folder, you'll find helpful types and interfaces. Keep in mind that an interface could be implemented as follows:
   ```js
   class Something implements MyInterface {}
   ```
9. A class named `Loader` is provided in your code, implementing a three-dot loader that appears in the gifs before any animation. While you can implement it, you can ignore it too and just directly trigger the required animations. Its main purpose is to serve as an example of how to implement a console animation.
10. Note that the loader is not a genuine loader; it merely waits for 3 seconds before displaying the selected easter egg.
