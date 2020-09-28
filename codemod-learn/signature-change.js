// jscodeshift -t signature-change.js signature-change.input.js -p
export default (fileInfo, api) => {
  const j = api.jscodeshift;
  const root = j(fileInfo.source);

  const ImportDeclaration = root.find(j.ImportDeclaration, {
    source: { type: 'Literal', value: 'car' }
  });
  const localName =
    ImportDeclaration
    .find(j.Identifier)
    .get(0)
    .node.name;

    const argKeys = [
      'color',
      'make',
      'model',
      'year',
      'miles',
      'bedliner',
      'alarm',
    ];

  return root.find(j.CallExpression, {
    callee: {
      object: { name: localName },
      property: { name: 'factory' }
    },
  }).replaceWith(nodePath => {
    const { node } = nodePath;
    const argumentsAsObject = j.objectExpression(
      node.arguments.map((arg, i) => 
        j.property(
          'init',
          j.identifier(argKeys[i]),
          j.literal(arg.value)
        )
      )
    );
    // 创建ast 对象
    // const object = j.objectExpression([
    //   j.property(
    //     'init',
    //     j.identifier('foo'),
    //     j.literal('bar')
    //   )
    // ]);
    node.arguments = [argumentsAsObject];
    return node;
  }).toSource();
}
