GTNH Modpack Zip File Structure:
```
.minecraft/
|-- config/
    |-- betterquesting/
        |-- DefaultQuests/
|-- mods/
|-- ...
```

Steam Age Patch Zip File Structure:
```
|-- mods/                        # Assembled from the mod forks in the GTNH-SteamAge GitHub Organiztion
|-- config/                      # Assembled from the GT-New-Horizons-Modpack repo fork, only a diff of changes
    |-- betterquesting/
        |-- DefaultQuests/
|-- README.md                    # Instructions on how to apply the Steam Age patch to the GTNH Modpack
|-- Steam Age vX.Y.Z Changelog   # (Rough) Summary of changes in this Steam Age version
|-- patch.sh                     # Shell patch script to apply to a pack version
|-- patch.bat                    # Windows patch script to apply to a pack version
```
