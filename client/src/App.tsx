import { AnimatePresence, motion } from "motion/react";
import React, { useState } from "react";
import "./App.css";

interface Todo {
  id: number;
  text: string;
  completed: boolean;
}

const App: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [input, setInput] = useState("");

  const addTodo = () => {
    if (input.trim() === "") return;
    const newTodo: Todo = {
      id: Date.now(),
      text: input,
      completed: false,
    };
    setTodos([...todos, newTodo]);
    setInput("");
  };

  const toggleComplete = (id: number) => {
    setTodos(
      todos.map((todo) =>
        todo.id === id ? { ...todo, completed: !todo.completed } : todo
      )
    );
  };

  const deleteTodo = (id: number) => {
    setTodos(todos.filter((todo) => todo.id !== id));
  };

  return (
    <div className="app">
      <h1 className="title">Todo List</h1>
      <div className="input-container">
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Add a new task..."
          className="p-2 mt-10 mr-2 text-lg font-semibold"
        />
        <button onClick={addTodo} className="px-3">
          Add
        </button>
      </div>
      <ul className="todo-list">
        <AnimatePresence>
          {todos.map((todo) => (
            <motion.li
              key={todo.id}
              initial={{ opacity: 0, x: -50 }}
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: 50 }}
              className={`todo-item ${todo.completed ? "completed" : ""}`}
            >
              <span onClick={() => toggleComplete(todo.id)}>{todo.text}</span>
              <button onClick={() => deleteTodo(todo.id)}>Delete</button>
            </motion.li>
          ))}
        </AnimatePresence>
      </ul>
    </div>
  );
};

export default App;
