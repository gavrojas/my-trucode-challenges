/** @type {import('tailwindcss').Config} */
import tailwindScrollbar from 'tailwind-scrollbar'
export default {
  content: [
    './src/App.vue',
    './src/components/**/*.{vue, js, ts, jsx, tsx}',
    './src/assets/**/*.css',
  ],
  theme: {
    extend: {
      colors:{
        "blue":'#2128a6',
        "toolbox":'#6f73bf',
        "lavender":'#dce0f2',
        "turquoise":'#30d9c8',
        "aqua":'#77d9cf',
        "white":'#ffffff',
        "black25":'rgba(0, 0, 0, 0.25)',
      },
      fontFamily: {
        sans: ['Roboto Condensed', 'sans-serif'],
      }, 
      spacing: {
        base: '1rem',
        sm: 'calc(1rem / 1.2)',
        lg: 'calc(1rem * 1.2)',
      },
      borderRadius: {
        base: '5px',
        xlg: 'calc(5px * 3.2)',
        lg: 'calc(5px * 1.2)',
        sm: 'calc(5px / 1.2)',
      },
      fontSize: {
        base: '1.2rem',
        h6: '1.2rem',
        h5: 'calc(1.2rem * 1.2)',
        h4: 'calc(1.2rem * 1.2 * 1.2)',
        h3: 'calc(1.2rem * 1.2 * 1.2 * 1.2)',
        h2: 'calc(1.2rem * 1.2 * 1.2 * 1.2 * 1.2)',
        h1: 'calc(1.2rem * 1.2 * 1.2 * 1.2 * 1.2 * 1.2)',
      },
    },
    plugins: [
      tailwindScrollbar,
    ],   
  },
}
