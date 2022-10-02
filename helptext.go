package main

var helpText = `
General options:
	-O   --output-document=FILE                 specify the file name
	-P   --directory-prefix=PREFIX              save files to PREFIX/...

HTTP options:
	--http-user=USER                            set http user to USER.
	--http-password=PASS                        set http password to PASS
	--no-cache                                  disallow server-cached data

HTTPS (SSL/TLS) options:
	--secure-protocol=PR                        choose secure protocol, one of auto, SSLv2,
	                                            SSLv3, TLSv1, TLSv1_1, TLSv1_2 and PFS
	--https-only                                 only follow secure HTTPS links
	--no-check-certificate                       don't validate the server's certificate
	--certificate=FILE                          client certificate file
	--certificate-type=TYPE                     client certificate type, PEM or DER
	--private-key=FILE                          private key file
	--private-key-type=TYPE                     private key type, PEM or DER
	--ca-certificate=FILE                       file with the bundle of CAs
	--ca-directory=DIR                          directory where hash list of CAs is stored
	--crl-file=FILE                             file with bundle of CRLs
	--pinnedpubkey=FILE/HASHES                  Public key (PEM/DER) file, or any number
	                                            of base64 encoded sha256 hashes preceded by
	                                            'sha256//' and separated by ';', to verify
	                                            peer against
	--random-file=FILE                          file with random data for seeding the SSL PRNG

	--ciphers=STR                               Set the priority string (GnuTLS) or cipher list string (OpenSSL) directly.
	                                            Use with care. This option overrides --secure-protocol.
	                                            The format and syntax of this string depend on the specific SSL/TLS engine.
FTP options:
	--ftp-user=USER                             set ftp user to USER
	--ftp-password=PASS                         set ftp password to PASS
	--no-remove-listing                         don't remove '.listing' files
	--no-glob                                   turn off FTP file name globbing
	--no-passive-ftp                            disable the "passive" transfer mode
	--preserve-permissions                      preserve remote file permissions
	--retr-symlinks                             when recursing, get linked-to files (not dir)
`