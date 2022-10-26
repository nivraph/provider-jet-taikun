import os
import yaml
import mdutils
from mdutils.mdutils import MdUtils

CRDS_DIR = "package/crds"
DOCS_DIR = "docs/"
TPL_DIR = "docs/templates/"


def open_yaml(filename):
    with open(file=filename, mode='r') as stream:
        try:
            data = yaml.safe_load(stream)
            return data
        except yaml.YAMLError as exc:
            print(exc)
            exit(1)


def readfile(filename):
    with open(filename, 'r') as file:
        return file.read()


def contains(l, elt):
    for e in l:
        if e == elt:
            return True
    return False


def is_field_present(block, fieldName):
    try:
        _ = block[fieldName]
        return True
    except KeyError:
        return False


def recursive_read(mdFile, block, field):
    if is_field_present(block[field], 'properties'):
        required = is_field_present(block[field], 'required')
        if required:
            for f in block[field]['properties']:
                if contains(block[field]['required'], f):
                    mdFile.write("\n* `"+f+"`: " +
                                 block[field]['properties'][f]['description'])
                    mdFile.write(" (Required)", color='red')

                    recursive_read(mdFile, block[field]['properties'], f)

                    mdFile.new_line()

        for f in block[field]['properties']:
            if not required or not contains(block[field]['required'], f):
                mdFile.write("\n* `"+f+"`: " +
                             block[field]['properties'][f]['description'])
                recursive_read(mdFile, block[field]['properties'], f)

                mdFile.new_line()
    else:
        return


def createMdDocs(data, filename):
    mdFile = MdUtils(file_name=DOCS_DIR+filename, title=filename)

    mdFile.new_paragraph("This document has been generated.\n")
    mdFile.new_line()

    mdFile.new_header(1, "Example")
    mdFile.insert_code(readfile(TPL_DIR+filename), language="yaml")
    mdFile.new_line()

    mdFile.new_header(1, "Schema")
    mdFile.new_line()

    required = True
    try:
        if len(data['required']) != 0:
            mdFile.new_header(2, "Required")
            mdFile.new_line()

            for field in data['properties']:
                if contains(data['required'], field):
                    mdFile.write("`"+field+"`: " +
                                 data['properties'][field]['description']+"\n")

                    recursive_read(mdFile, data['properties'], field)

                    mdFile.new_line()
    except KeyError:
        required = False
        pass

    mdFile.new_header(2, "Optional")
    mdFile.new_line()

    for field in data['properties']:
        if not required or not contains(data['required'], field):
            mdFile.write("`"+field+"`: " +
                         data['properties'][field]['description']+"\n")

            recursive_read(mdFile, data['properties'], field)

            mdFile.new_line()

    mdFile.create_md_file()


def generate_resource_docs(filename):
    print("generating docs for: ", filename)
    data = open_yaml(filename)

    resource_data = data['spec']['versions'][0]['schema']['openAPIV3Schema'][
        'properties']['spec']['properties']['forProvider']

    mdfilename = data['spec']['group']
    createMdDocs(resource_data, mdfilename)


def generate_provider_docs(filename):
    print("generating docs for: ", filename)
    data = open_yaml(filename)

    return

#
# --------------------------------------------------------------------------------------------------------------

# --------------------------------------------------------------------------------------------------------------
#


if not os.path.exists(path=CRDS_DIR):
    print("CRD's do not exists.")


for filename in os.listdir(CRDS_DIR):
    f = os.path.join(CRDS_DIR, filename)
    if os.path.isfile(f) and (filename.endswith('.yaml') or filename.endswith('.yml')) and not filename.startswith('taikun.'):
        generate_resource_docs(f)
    elif os.path.isfile(f) and (filename.endswith('.yaml') or filename.endswith('.yml')) and filename.startswith('taikun.'):
        generate_provider_docs(f)
    else:
        continue
