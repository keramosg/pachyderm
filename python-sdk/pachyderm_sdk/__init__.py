from .client import Client
from .batch_datums import batch_datums

# Python version compatibility.
try:
    # >= 3.8
    import importlib.metadata as metadata
except ImportError:
    #  < 3.8
    import importlib_metadata as metadata  # type: ignore


__pdoc__ = {"proto": False}

__all__ = [
    "Client",
    "batch_datums"
]

__version__ = ""
try:
    __version__ = metadata.version(__name__)  # type: ignore
except (FileNotFoundError, ModuleNotFoundError):
    pass


from .api.pfs import _additions as __pfs_additions
