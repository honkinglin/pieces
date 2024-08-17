import { VirtualList } from "./virtual-list.js";
import { initMockDB } from "./utils/db.js";

const container = document.body;
const DB = initMockDB({
  title: "Frontend System Design",
  body: "Frontend System Design",
});
const template = document.getElementById("card_template");

const updateTemplate = (element, datum) => {
  const [cardTitle] = element.getElementsByTagName("h3");
  const [cardBody] = element.getElementsByTagName("section");
  [cardTitle.textContent, cardBody.textContent] = [datum.title, datum.body];
};

const getTemplate = (datum) => {
  const element = template.content.firstElementChild.cloneNode(true);
  updateTemplate(element, datum);
  return element;
};

const list = new VirtualList(container, {
  getPage: (p) => DB.getPage(p),
  getTemplate,
  updateTemplate,
  pageSize: 10,
});
list.render();
