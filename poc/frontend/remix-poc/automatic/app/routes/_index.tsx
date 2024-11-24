import type { MetaFunction } from "@remix-run/node";
import { Link } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "To Do App" }
  ];
};

export default function Index() {
  return (
    <div className="flex flex-col justify-center items-center gap-48 h-[70vh]">
      <div className="">
        <h1 className="text-[10rem] font-semibold">To Do List</h1>
      </div>
      <div className="">
        <Link to="/todo">
          <button className="bg-white text-black text-7xl py-4 px-6 rounded-lg">New List</button>
        </Link>
      </div>
    </div>
  );
}
