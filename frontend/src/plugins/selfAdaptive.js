const adjustFontSize = () => {
  const root = document.documentElement;
  const w = root.clientWidth;
  let fz = w / 10;
  fz = fz.toFixed(1);
  fz = Math.max(fz, 15);
  fz = Math.min(fz, 54);
  root.style.fontSize = `${fz}px`;
}

export default {
  adjustFontSize
}