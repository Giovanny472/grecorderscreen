
#define MyAppName "gScreenShot"
#define MyAppVersion "0.1"
#define MyAppPublisher "@egmn472"
#define MyAppURL "https://www.github.com/Giovanny472"
#define MyAppExeName "gscreenshot.exe"
[Setup]
AppId={{316F5133-94BC-45E7-A57D-95B09460F757}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
AppVerName={#MyAppName} {#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DefaultDirName={autopf}\gscreenshot
DisableProgramGroupPage=yes
;LicenseFile=Free
; Uncomment the following line to run in non administrative install mode (install for current user only.)
;PrivilegesRequired=lowest
OutputDir=..\setup\
OutputBaseFilename=gScreenshot_setup
SetupIconFile=..\icon\gscreenshot.ico
Compression=lzma
SolidCompression=yes
WizardStyle=modern
[Languages]
Name: "russian"; MessagesFile: "compiler:Languages\Russian.isl"
[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked
[Files]
Source: "..\..\bin\{#MyAppExeName}"; DestDir: "{app}"; Flags: ignoreversion
[Icons]
Name: "{autoprograms}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; HotKey: "ctrl+alt+s"
Name: "{autodesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon
[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall skipifsilent