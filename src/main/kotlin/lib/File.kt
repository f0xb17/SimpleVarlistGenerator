/**
 * libfile.kt
 * Extracts all variables from an OMSI script file (.osc) and saves them in a newly created varlist or stringvarlist.
 *
 * Created by f0xb17 on 04/21/2024.
 * Copyright © 2024 f0xb17. All rights reserved.
 */

package org.foxbit.lib

import java.io.File

/**
 * This class provides all possible functions to automatically generate a varlist or stringvarlist.
 * @param [dir] The file path that contains the .osc files.
 * @param [option] 1 = varlist / 2 = stringvarlist
 */
class File(private val dir: String, private val option: Int) {

    /**
     * This method is used to find all files with the extension .osc at the specified storage point
     * and save them in a list.
     * @param [suffix] The File extension to be searched for. (e.g. .osc)
     * @return a List of Strings, which contains all files with the ending .osc in a specific folder.
     */
    private fun findFilesInFolder(suffix: String): List<String> {
        val files = File(this.dir).walk()
            .filter { it.isFile && it.name.endsWith(suffix) }
                .toList()
                    .map { it.absolutePath }
        return files
    }

    /**
     * This method is used to save the content of an .osc file into a searchable list of strings.
     * @param [filePath] Represents a file at the specified file path.
     * @return a List of Strings, which contains every line of @link[filePath]
     */
    private fun readFileAsLines(filePath: String): List<String> {
        return File(filePath).readLines().map { it.lowercase() }
    }

}