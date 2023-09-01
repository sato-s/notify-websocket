'use client';

import React from "react";

export default function Room({params}: {params: {slug: string}}) {
  return <div>My Post: {params.slug}</div>
}
