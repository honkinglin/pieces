import { createElement, render, Component } from './toy-react';

class MyComponent extends Component {
  constructor() {
    super();
    this.state = {
      a: 'test',
    };
  }

  render() {
    return (
      <div>
        <span>{this.state.a}</span>
        {this.children}
      </div>
    );
  }
}

render(<MyComponent id="a">
  <span>123123</span>
  <span></span>
  <span></span>
</MyComponent>, document.body);
