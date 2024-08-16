// src/components/UptimeTable.js
import React from 'react';

function UptimeTable({ data }) {
    return (
        <table>
            <thead>
                <tr>
                    <th>URL</th>
                    <th>Status</th>
                    <th>Checked At</th>
                </tr>
            </thead>
            <tbody>
                {data.map((item, index) => (
                    <tr key={index}>
                        <td>{item.url}</td>
                        <td>{item.status ? 'Online' : 'Offline'}</td>
                        <td>{item.checked_at}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
}

export default UptimeTable;
