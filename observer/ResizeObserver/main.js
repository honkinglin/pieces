/*
 * 1. Initiate Resize observer
 * 2. Implement callback - use inlineSize and blockSize of the entries for width and height
 * 3. If width and height is less than 150px
 *   3.1 - Apply border-radius: 100%
 *   3.2 - Apply border-width: 4px
 * 4. If width and height is more than 150px = unset the properties above
 * 5. All 4 boxes should use same observer
 */
const observer = new ResizeObserver((entries) => {
  for (const entry of entries) {
    const target = entry.target;
    const box = entry.borderBoxSize[0];
    if (box.blockSize < 150 && box.inlineSize < 150) {
      target.style.borderRadius = "100%";
      target.style.borderWidth = "4px";
    } else {
      target.style.borderRadius = "unset";
      target.style.borderWidth = "unset";
    }
  }
});

const boxes = document.querySelectorAll(".box");
boxes.forEach((box) => observer.observe(box));
