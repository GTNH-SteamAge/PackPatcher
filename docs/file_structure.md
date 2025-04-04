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
|-- steamagemods/                 # Assembled from the mod forks in the GTNH-SteamAge GitHub Organiztion
|-- steamageconfig/               # Assembled from the GT-New-Horizons-Modpack repo fork, only a diff of changes
    |-- betterquesting/
        |-- DefaultQuests/
|-- Steam_Age_README.md           # Instructions on how to apply the Steam Age patch to the GTNH Modpack
|-- Steam_Age_vX.Y.Z_Changelog.md # (Rough) Summary of changes in this Steam Age version
|-- steam_age_patch.sh            # Shell patch script to apply to a pack version
|-- steam_age_patch.bat           # Windows patch script to apply to a pack version
```
