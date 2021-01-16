
import React from 'react';
const axios = require('axios').default;
class MyComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: [],
      value: ''
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleChange(event) {    this.setState({value: event.target.value});  }
  handleSubmit(event) {
    console.log(this.state.value)
    if (!this.state.value){
      this.props.onChange(null)
    }

    //console.log(typeof Number(this.state.value))
    event.preventDefault();
    axios({
  method: 'post',
  url: '/articleval',
  data: {
    "Ccal":Number(this.state.value)
}
}).then((response) => {

  console.log(response['data']);
  this.setState({
    isLoaded: true,
    items: response['data']
  });
}, (error) => {
  console.log(error);
});
}


  render() {
    const { error, isLoaded, items } = this.state;
    if (error) {
      return (
      <form onSubmit={this.handleSubmit} name="frm"> <label>
      Ccal:
      <input type="number"  value={this.state.value} onChange={this.handleChange} onclick="return IsEmpty();"/>        </label>
      <input type="submit" value="Submit" />
      </form>);
      //<div>Ошибка: {error.message}</div>;
    }  else {
      return (
        <ul>
        <form onSubmit={this.handleSubmit} name="frm"> <label>
        Ccal:
        <input type="number" value={this.state.value} onChange={this.handleChange} onclick="return IsEmpty();"/>        </label>
        <input type="submit" value="Submit" />
        </form>
          {items.map(item => (
            <li key={item.Id}>
              {item.Name} {item.Ccal} 100гр
            </li>
          ))}
        </ul>
     );
    }
  }
}
export default MyComponent;
//ReactDOM.render(<MyComponent />, document.getElementById('root'));
