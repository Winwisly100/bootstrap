# i18n bottom up

We have modules and then maintemplate.

We want the arb files from each module to have all its translations, and then for maintemplate to be able to use them.


https://github.com/flutter/flutter/tree/master/dev/tools/localization

https://pascalw.me/blog/2020/03/09/localizing-your-app-with-flutters-new-gen-l10n-tool.html


gen_l10n.dart
- looks in l10n folder
- looks for app_en.arb file
- outputs app_localizations.dart 
- outputs AppLocalizations.dart