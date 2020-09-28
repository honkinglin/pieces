// jscodeshift deprecated.js deprecated.input.js

export default (fileInfo, api) => {
  const j = api.jscodeshift;
  const root = j(fileInfo.source);

  const importDeclaration = root.find(j.ImportDeclaration, {
    source: {
      type: 'Literal',
      value: 'geometry',
    }
  });

  // get the local name for the imported module
  const localName =
    importDeclaration
    .find(j.Identifier)
    .get(0)
    .node.name;

  return root.find(j.MemberExpression, {
    object: {
      name: localName,
    },
    property: {
      name: 'circleArea',
    },
  })
  .replaceWith(nodePath => {
    // get the underlying Node
    const { node } = nodePath;
    // change to our new prop
    node.property.name = 'getCircleArea';
    // replaceWith should return a Node, not a NodePath
    return node;
  })
  .toSource();
}