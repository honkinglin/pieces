import { createCardElement, getHeading } from "./utils.js";
import { initMockDB } from "./db.js";

const SUPPORTED_ELEMENTS = new Set(["/h1", "/h2", "/h3"]);

const db = initMockDB({
  title: "Fundamentals of Frontend System Design",
  body: "Learning to use Intersection Observer",
});
const list = document.getElementById("list");
const observerElement = document.getElementById("bottom-observer");

const mutationObserver = new MutationObserver((entries) => {
  /*
   * @todo
   *
   * 1. Loop through mutations (first arg of callback)
   * 2. use mutation.type === 'characterData'
   * 3. When the content matches with any supported tag in SUPPORTED_ELEMENTS
   *    use getHeading function to convert text to HTMLHeading element
   * 4. Replace the node with a newly created node
   * 5. Keep focus on the element
   */
  for (const mut of entries) {
    if (mut.type === "characterData") {
      const target = mut.target;
      if (
        mut.type === "characterData" &&
        SUPPORTED_ELEMENTS.has(target.textContent)
      ) {
        const heading = getHeading(target);
        target.replaceWith(heading);
        heading.focus();
      }
    }
  }
});
let page = 0;
const observer = new IntersectionObserver(
  async ([bottom]) => {
    if (bottom.isIntersecting) {
      observerElement.textContent = "Loading";
      const data = await db.getPage(page++);
      const fragment = new DocumentFragment();
      data.forEach(({ title, body }) => {
        const card = createCardElement(title, body);
        fragment.appendChild(card);
        // @todo register observer
        mutationObserver.observe(card, { subtree: true, characterData: true });
      });

      list.appendChild(fragment);
    }
  },
  { threshold: 0.1 }
);
observer.observe(observerElement);
