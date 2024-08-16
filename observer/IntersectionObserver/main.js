import { initMockDB } from "./db.js";

const [template, list, observerElement] = document.querySelectorAll(
  "#card_template, #list, #bottom-observer"
);

const db = initMockDB({
  title: "Fundamentals of Frontend System Design",
  body: "Learning to use Intersection Observer",
});

function createCardElement(title, body) {
  const element = template.content.cloneNode(true).firstElementChild;
  const [cardTitle, cardBody] = element.querySelectorAll(
    ".card__title, .card__body__content"
  );
  [cardTitle.textContent, cardBody.textContent] = [title, body];
  return element;
}

/**
 * Exercise - Intersection Observer
 * 1. Create Intersection observer instance and provide a callback to it
 * 2. In the callback use mock db - next function to get the next chunk of data
 * 3. Create a fragment where you chunk all your DOM Mutations
 * 4. Update fragment
 * 5. Append fragment to "list" container
 */
let page = 0;
const observer = new IntersectionObserver(
  async ([entry]) => {
    if (entry.isIntersecting) {
      const data = await db.getPage(page++);
      const fragment = new DocumentFragment();
      for (const datum of data) {
        const card = createCardElement(datum.title, datum.body);
        fragment.appendChild(card);
      }
      list.appendChild(fragment);
    }
  },
  { threshold: 0.2 }
);

observer.observe(observerElement);
