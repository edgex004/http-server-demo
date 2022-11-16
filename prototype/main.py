from typing import Union

from fastapi import FastAPI, Query
from pydantic import BaseModel

import interface_types

description = """
Test server to store and retrieve log entries from volitile memory.

## Log entries

You can:
* **create log entries**.
* **read log entries**.
"""

tags_metadata = [
    {
        "name": "log entries",
        "description": "Manage log entries.",
    },
]

app = FastAPI(
    version="0.0.2",
    description=description,
    openapi_tags=tags_metadata
    )

Logs: dict[str,list[interface_types.LogEntry]] = {}

@app.get("/logs/entries", \
    response_model=list[interface_types.GetLogEntryResponse], \
    tags=["log entries"])
def read_entries(\
    entry_ids: Union[list[int], None] = Query(default=None), \
    device_ids: Union[list[str], None] = Query(default=None)):
    """
    Read existing log entries.
    :param entry_ids: 0 based indexes for desired log entries 
    :param device_ids: strings that must match the IDs of the desired robots
    :return: the query result
    """
    ret = []
    if device_ids:
        for device_id in device_ids:
            if device_id in Logs and entry_ids:
                for entry_id in entry_ids:
                    if entry_id >= 0 and entry_id < len(Logs[device_id]):
                        ret.append({"entry_id": entry_id, "entry":Logs[device_id][entry_id]})
    return ret

@app.post("/logs/entries", \
    response_model=interface_types.PostLogEntryResponse, \
    tags=["log entries"])
def add_entry(entry: interface_types.LogEntry):
    """
    Add a new log entry.
    :param entry: 0 based indexes for desired log entries 
    :return: the id for the new entry. This will increment for each entry added for a given robot.
    """
    if entry.device_id not in Logs:
        Logs[entry.device_id] = []
    Logs[entry.device_id].append(entry)
    entry_id = len(Logs[entry.device_id]) - 1
    return {"entry_id": entry_id}