import React from 'react';

export default function Navigation({ results }) {
  if (results === undefined || results.length === 0) {
    return <div />;
  }
  const entries = results.map((result, index) => {
    var subMe;
    subMe = (
      <ul>
        <Navigation results={result.subs} />
      </ul>
    );

    return (
      <li>
        {result.name}
        {subMe}
      </li>
    );
  });

  return entries;
}
