import { useState, ChangeEvent } from "react";
import { Link } from "@remix-run/react";

export default function TodoPage() {
  const [tasks, setTasks] = useState<string[]>([]);
  const [newTask, setNewTask] = useState<string>("");

  const addTask = () => {
    if (newTask.trim()) {
      setTasks((prevTasks) => [...prevTasks, newTask]);
      setNewTask("");
    }
  };

  const removeTask = (index: number) => {
    setTasks((prevTasks) => prevTasks.filter((_, i) => i !== index));
  };

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    setNewTask(e.target.value);
  };

  return (
    <div>
      <Link to="/">
        <button>back</button>
      </Link>
      <div className="flex flex-col items-center h-screen">
        <div>
          <h1>To-Do List</h1>
        </div>
        <div>
          <div>
            <input
              type="text"
              value={newTask}
              onChange={handleInputChange}
              placeholder="Enter a new task"
              className="py-1"
            />
            <button onClick={addTask} style={{ padding: "5px 10px" }}>
              Add
            </button>
          </div>
          <div>
            <ul style={{ marginTop: "20px", listStyle: "none", padding: "0" }}>
              {tasks.map((task, index) => (
                <li
                  key={index}
                  style={{
                    marginBottom: "10px",
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                  }}
                >
                  {task}
                  <button
                    onClick={() => removeTask(index)}
                    style={{
                      marginLeft: "10px",
                      padding: "2px 5px",
                      background: "red",
                      color: "white",
                      border: "none",
                      cursor: "pointer",
                    }}
                  >
                    Delete
                  </button>
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
