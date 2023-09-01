'use client';

import React from "react";
import Room from "../lib/Room";

export default function Page({params}: {params: {slug: string}}) {
  const room = new Room('ss');
  return <div>My Post: {params.slug}</div>
}
