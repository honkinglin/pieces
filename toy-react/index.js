import { createElement, render, Component } from './toy-react';

class MyComponent extends Component {
  constructor() {
    super();
    this.state = {
      a: '1',
    };
  }

  render() {
    return (
      <div>
        <span>{this.state.a}</span>
        <button onClick={() => { this.state.a ++; this.rerender(); } }>click</button>
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
