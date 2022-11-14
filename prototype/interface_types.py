from pydantic import BaseModel
from typing import Union

class Location(BaseModel):
    latitude: float
    longitude: float

class LogEntry(BaseModel):
    device_id: str
    title: str
    entry: str
    location: Union[Location, None]
    temperature: Union[float, None]


class GetLogEntryResponse(BaseModel):
    entry_id: int
    entry: LogEntry

class PostLogEntryResponse(BaseModel):
    entry_id: int